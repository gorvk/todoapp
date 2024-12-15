import { CommonResponse, TodoData } from "../models/types";
import { getMethodDeleteHeader, getMethodGetHeader, getMethodPostHeader, getMethodPutHeader } from "../utils/https";

const baseUrl = 'http://localhost:9090/api/todo';

export async function getAllTodos(): Promise<TodoData[]> {
    const headers = getMethodGetHeader();
    const response = await fetch(baseUrl, headers);
    const data: TodoData[] = await response.json()
    return data;
}

export async function addTodo(todo: TodoData): Promise<CommonResponse> {
    const headers = getMethodPostHeader(todo)
    const response = await fetch(baseUrl, headers);
    const data: CommonResponse = await response.json()
    return data;
}

export async function updateTodo(todo: TodoData): Promise<CommonResponse> {
    const headers = getMethodPutHeader(todo)
    const response = await fetch(baseUrl, headers);
    const data: CommonResponse = await response.json()
    return data;
}

export async function deleteTodo(id: string): Promise<CommonResponse> {
    const url = baseUrl + `/${id}`;
    const headers = getMethodDeleteHeader()
    const response = await fetch(url, headers);
    const data: CommonResponse = await response.json()
    return data;
}