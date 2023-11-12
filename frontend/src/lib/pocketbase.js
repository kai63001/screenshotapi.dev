import Pocketbase from 'pocketbase';

import { writable } from 'svelte/store';

export const pb = new Pocketbase('http://127.0.0.1:8090/');

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(async (auth) => {
    // console.log('auth changed', auth)
    // get access key from table access_keys
    // pb.authStore.model.access_key
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