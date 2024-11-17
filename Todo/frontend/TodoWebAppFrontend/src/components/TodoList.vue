<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useTodoStore } from '@/stores/todo';
import type { Todo } from '@/models/todo';

const store = useTodoStore();
const newTodo = ref('');
const editingDescriptions = ref<Record<number, string>>({});
onMounted(async () => {
    await store.fetchTodos();
    store.todos.forEach(todo => {
        editingDescriptions.value[todo.id] = todo.description
    });
});

async function updateTodo(todo: Todo) {
    console.log(editingDescriptions.value[todo.id])
    const oldDescrtiption = todo.description;
    todo.description = editingDescriptions.value[todo.id];
    if (todo.description !== oldDescrtiption) {
        await store.changeTodoStatus(todo, false);
    }

    todo.isEditing = false;
}

async function deleteTodo(todo: Todo){
    await store.deleteTodo(todo)
}


</script>

<template>
    <div>
        <h1 id="subtitle">Todo List</h1>
        <p v-if="store.apiError" class="error">{{ store.apiError }}</p>
        <div v-else="store.apiError">

            <div class="form-container">
                <form @submit.prevent="store.addTodo(newTodo)" class="todo-new-form">
                    <input type="text" v-model="newTodo" required placeholder="new todo">
                    <button>Add Todo</button>
                </form>
            </div>
            <ul>
                <li v-for="todo in store.todos" :key="todo.id" class="todo-item">
                    <div v-if="todo.isEditing">
                        <input type="text" v-model="editingDescriptions[todo.id]" />
                    </div>
                    <div v-else="todo.isEditing">
                        {{ todo.description }}
                    </div>

                    <div class="input-container">

                        <button v-if="todo.isEditing" @click="updateTodo(todo)">Save</button>
                        <button @click="todo.isEditing = !todo.isEditing" class="item-btn">Edit</button>
                        <button @click="deleteTodo(todo)" class="item-btn">Delete</button>
                        <input type="checkbox" :checked="todo.isDone" @change="() => store.changeTodoStatus(todo, true)"
                            class="checkbox" />
                    </div>
                </li>
            </ul>
        </div>
    </div>
</template>

<style>
.done {
    text-decoration: line-through;
}

form {
    display: flex;
    gap: 10px;
}

ul {
    list-style: none;
    padding: 10px;
}

#subtitle {
    text-align: center;
    padding: 30px 0 30px 0;
}

input[type="checkbox"] {
    cursor: pointer;
}

input[type="text"] {
    padding: 5px;
    border-radius: 5px;
    border: 1px solid #ccc;
    font-size: 16px;
    flex-grow: 0.5;
}

button {
    padding: 5px;
    border-radius: 5px;
    border: 1px solid #ccc;
    font-size: 16px;
    background-color: #f0f0f0;
    cursor: pointer;
}

button:hover {
    background-color: #e0e0e0;
}

form-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100hv;
}

.todo-new-form {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 60px;
}

.todo-item {
    font-size: 20px;
    font-weight: 700;
    max-width: 70%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
    border-bottom: 1px solid #ccc;
    margin: 0 auto;
}

.input-container {
    display: flex;
    align-items: center;
    column-gap: 10px;
}
</style>