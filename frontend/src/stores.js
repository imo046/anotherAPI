import { writable } from 'svelte/store';

export const entries = writable([]);

const apiBase = 'http://localhost:4040';

export const fetchEntries = async () => {
    const res = await fetch(`${apiBase}/getAll`);
    const data = await res.json();
    entries.set(data);
}

export const createEntry = async (entry) => {
    const res = await fetch(`${apiBase}/create`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(entry),
    });
    await fetchEntries();
}

export const updateEntry = async (entry) => {
    const { id, value } = entry;
    const res = await fetch(`${apiBase}/update/${id}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ value }),
    });
    await fetchEntries();
}

export const deleteEntry = async (id) => {
    const res = await fetch(`${apiBase}/delete/${id}`, {
        method: 'DELETE',
    });
    await fetchEntries();
}

export const deleteAllEntries = async () => {
    const res = await fetch(`${apiBase}/deleteAll`, {
        method: 'DELETE',
    });
    await fetchEntries();
}
