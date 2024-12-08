import { CommonResponse, TodoData } from "../models/types";

const baseUrl = 'http://localhost:9090/api/todo';

export async function getAllTodos(): Promise<TodoData[]> {
    const response = await fetch(baseUrl);
    const data: TodoData[] = await response.json()
    return data;
}

export async function addTodo(todo: TodoData): Promise<CommonResponse> {
    const requestInfo: RequestInfo = new Request(baseUrl, {
        method: 'POST',
        body: JSON.stringify(todo),
    })
    const response = await fetch(requestInfo);
    const data: CommonResponse = await response.json()
    return data;
}

export async function updateTodo(todo: TodoData): Promise<CommonResponse> {
    const requestInfo: RequestInfo = new Request(baseUrl, {
        method: 'PUT',
        body: JSON.stringify(todo),
    })
    const response = await fetch(requestInfo);
    const data: CommonResponse = await response.json()
    return data;
}

export async function deleteTodo(id: string): Promise<CommonResponse> {
    const requestInfo: RequestInfo = new Request(baseUrl + `/${id}`, { method: 'DELETE' })
    const response = await fetch(requestInfo);
    const data: CommonResponse = await response.json()
    return data;
}