import axios from 'axios';
import Pocketbase from 'pocketbase';
import { writable } from 'svelte/store';

export const pb = new Pocketbase(import.meta.env.VITE_API_URL || 'http://127.0.0.1:8090/');

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(async (auth) => {
    // console.log(auth)
    if (!auth) {
        currentUser.set(null)
        localStorage.removeItem('access_key')
        return
    }
    const record = await pb.collection('access_keys').getFirstListItem(`user_id="${pb.authStore.model?.id}"`)
    const model = {
        ...pb.authStore.model,
    }
    localStorage.setItem('access_key', record.access_key)
    currentUser.set(model)
})


export const axiosInstance = axios.create({
    baseURL: `${import.meta.env.VITE_API_URL}/v1`,
    headers: {
        Authorization: 'Bearer ' + pb?.authStore?.token
    }
});