import React, { useState } from 'react';
import { useMutation } from '@apollo/client';
import gql from 'graphql-tag';

const CREATE_TODO = gql`
    mutation createTodo($title: String!) {
        createTodo(title: $title) {
            id
            title
            completed
        }
    }
`;

function CreateTodo() {
    const [title, setTitle] = useState('');
    const [createTodo, { error }] = useMutation(CREATE_TODO, {
        variables: {
            title,
        },
        update(cache, { data: { createTodo } }) {
            cache.modify({
                fields: {
                    todos(existingTodos = []) {
                        const newTodoRef = cache.writeFragment({
                            data: createTodo,
                            fragment: gql`
                                fragment NewTodo on Todo {
                                    id
                                    title
                                    completed
                                }
                            `,
                        });
                        return [...existingTodos, newTodoRef];
                    },
                },
            });
        },
    });

    const handleSubmit = (e) => {
        e.preventDefault();
        createTodo({
            variables: {
                title,
            },
        });
        setTitle('');
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                type="text"
                placeholder="New todo"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
            />
        <button type="submit">Create Todo</button>
        </form>
    );
}

export default CreateTodo;
