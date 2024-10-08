// place files you want to import through the `$lib` alias in this folder.
import { writable } from 'svelte/store';

export interface Entry {
    id: string;
    entry_val: string;
}


const apiBase = 'http://localhost:4040';

export const fetchEntries = async () => {
    const res = await fetch(`${apiBase}/getAll`,{
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
        },


    });
    if (!res.ok) {
        throw new Error('Failed to fetch entries');
    }
    return await res.json();

}


export async function uploadEntries(file: File) {
    const formData = new FormData();
    formData.append("file", file);

    try {
        const response = await fetch(`${apiBase}/upload`, {
            method: "POST",
            body: formData,
        });

        if (!response.ok) {
            throw new Error("Failed to upload CSV.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error uploading CSV:", error);
        throw error; // Re-throw the error to handle it in the Svelte component
    }

}

export const updateEntry = async (entry: { id: string; value: string }) => {
    const { id, value } = entry;
    const response = await fetch(`${apiBase}/update/${id}`, {
        method: 'PATCH',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ value }),
    });
    if (!response.ok) {
        throw new Error('Failed to update entry');
    }
    console.log(response)

    return await response.json();
}

export const deleteEntry = async (id: string) => {
    const response = await fetch(`${apiBase}/delete/${id}`, {
        method: 'DELETE',
    });

    if (!response.ok) {
        throw new Error('Failed to delete entry');
    }

    console.log(response)
    return await response.json();
}

export const deleteAllEntries = async () => {
    await fetch(`${apiBase}/deleteAll`, {
        method: 'DELETE',
    });
    await fetchEntries();
}
