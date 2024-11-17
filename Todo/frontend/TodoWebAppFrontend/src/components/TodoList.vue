<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { useTodoStore } from '@/stores/todo';

const store = useTodoStore();
const newTodo = ref('');
onMounted(() => {
    store.fetchTodos();
});

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
                    {{ todo.description }}

                    <input type="checkbox" :checked="todo.isDone" @change="() => store.changeTodoStatus(todo, true)"
                        class="checkbox" />
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

.checkbox {
    margin-left: auto;
}
</style>