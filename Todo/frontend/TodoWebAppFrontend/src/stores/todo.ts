import { defineStore } from 'pinia'
import type { Todo } from '../models/todo'
import { TodoParser } from '../models/todo'

export const useTodoStore = defineStore('todo', {
  state: () => ({
    todos: [] as Todo[],
    apiError: null as string | null,
  }),
  getters: {},
  actions: {
    async addTodo(desc: string) {
      try {
        const response = await fetch(`http://localhost:8080/todos`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            description: desc,
          }),
        })
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        const rawData = await response.json()
        console.log(rawData)
        const newTodo = TodoParser.parseTodo(rawData)
        this.todos.push(newTodo)
      } catch (error) {
        this.apiError = (error as Error)?.message
      }
    },
    async fetchTodos() {
      try {
        const response = await fetch('http://localhost:8080/todos')
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        const rawData = await response.json()
        console.log(rawData)
        const todoList: Todo[] = TodoParser.parseTodos(rawData)
        this.todos = todoList
      } catch (error) {
        console.error('Error fetching todos', error)
        this.apiError = (error as Error)?.message
      }
    },

    async changeTodoStatus(todo: Todo, isCheck: boolean) {
      try {
        const response = await fetch(`http://localhost:8080/todos/${todo.id}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            id: todo.id,
            description: todo.description,
            isDone: isCheck ? !todo.isDone : todo.isDone,
          }),
        })
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        const rawData = await response.json()
        console.log(rawData)
        const updatedTodo = TodoParser.parseTodo(rawData)
        const index = this.todos.findIndex((t) => t.id === updatedTodo.id)
        this.todos[index] = updatedTodo
      } catch (error) {
        this.apiError = (error as Error)?.message
      }
    },
    async deleteTodo(todo: Todo) {
      try {
        const response = await fetch(`http://localhost:8080/todos/${todo.id}`, {
          method: 'DELETE',
        })
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`)
        }
        this.todos = this.todos.filter((t) => t.id !== todo.id)
      } catch (error) {
        this.apiError = (error as Error)?.message
      }
    },
  },
})
