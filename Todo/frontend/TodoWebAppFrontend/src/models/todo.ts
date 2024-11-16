export interface Todo {
    id: number;
    description: string;
    isDone: boolean;
}
export class TodoParser {
    static parseTodos(rawData: any[]): Todo[] {
        return rawData.map((todo) => {
            return {
                id: todo.id,
                description: todo.description,
                isDone: todo.isDone,
            };
        });
    }
    static parseTodo(rawData: any): Todo {
        return {
            id: rawData.id,
            description: rawData.description,
            isDone: rawData.isDone,
        };

    }
}

