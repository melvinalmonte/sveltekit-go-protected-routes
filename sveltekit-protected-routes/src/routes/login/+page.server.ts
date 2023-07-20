import type { Actions } from './$types';
import { redirect } from '@sveltejs/kit';

export const actions = {
	default: async (event) => {
		const formData = await event.request.formData();
		const req = await fetch('http://localhost:1323/api/login', {
			method: 'POST',
			body: formData
		});
		if (!req.ok) {
			throw redirect(301, '/login');
		}
		const resp = await req.json();

		event.cookies.set('session', resp.token);
		throw redirect(301, '/');
	}
} satisfies Actions;
