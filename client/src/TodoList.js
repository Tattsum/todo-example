import React from "react";
import { useQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";

const GET_TODOS = gql`
    query {
        todos {
            id
            text
            completed
        }
    }
`;

function TodoList() {
    const { loading, error, data } = useQuery(GET_TODOS);
    if (loading) return "Loading...";
    if (error) return `Error! ${error.message}`;

    return (
        <ul>
            {data.todos.map((todo) => (
                <li key={todo.id}>
                    {todo.text} - {todo.completed ? "Completed" : "Incomplete"}
                </li>
            ))}
        </ul>
    );
}

export default TodoList;
