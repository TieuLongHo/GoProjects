<script setup lang="ts">
import { onMounted, ref } from 'vue';
import type { Todo } from '../models/todo';
import { TodoParser } from '../models/todo';

const todos = ref<Todo[]>([]);
const apiError = ref<string | null>(null);
const newTodo = ref('');

async function addTodo(desc: string) {
    try {
        const response = await fetch(`http://localhost:8080/todos`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                description: desc,
            }),
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const rawData = await response.json();
        console.log(rawData);
        const newTodo = TodoParser.parseTodo(rawData);
        todos.value.push(newTodo)
    } catch (error) {
        apiError.value = (error as Error).message;
    }
}
async function fetchTodos() {
    try {
        const response = await fetch('http://localhost:8080/todos');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const rawData = await response.json();
        console.log(rawData);
        todos.value = TodoParser.parseTodos(rawData);
    } catch (error) {
        apiError.value = (error as Error).message;
    }
}

async function changeTodoStatus(todo: Todo, isCheck: boolean) {
    try {
        const response = await fetch(`http://localhost:8080/todos/${todo.id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                id: todo.id,
                description: todo.description,
                isDone: isCheck?!todo.isDone:todo.isDone,
            }),
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const rawData = await response.json();
        console.log(rawData);
        const updatedTodo = TodoParser.parseTodo(rawData);
        const index = todos.value.findIndex((t) => t.id === updatedTodo.id);
        todos.value[index] = updatedTodo;
    } catch (error) {
        apiError.value = (error as Error).message;
    }
}

onMounted(() => {
    fetchTodos();
});

</script>

<template>
    <div>
        <h1>Todo List</h1>
        <p v-if="apiError" class="error">{{ apiError }}</p>
        <form @submit.prevent="addTodo(newTodo)">
    <input v-model="newTodo" required placeholder="new todo">
    <button>Add Todo</button>
  </form>        <ul>
            <li v-for="todo in todos" :key="todo.id">
                <input type="checkbox" :checked="todo.isDone" @change="() => changeTodoStatus(todo,true)" />
                {{ todo.description }}
            </li>
        </ul>
    </div>
</template>


<style>
.done {
  text-decoration: line-through;
}
</style>