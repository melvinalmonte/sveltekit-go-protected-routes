import type { Handle } from '@sveltejs/kit';
import { redirect } from '@sveltejs/kit';

export const handle = (async ({ event, resolve }) => {
	// check if we have an auth token
	const authToken = event.cookies.get('session');

	// if we do, extract the payload
	const tokenPayload = authToken ? authToken.split('.')[1] : null;

	// if we have a payload, set the user in locals
	const decodedPayload = tokenPayload ? JSON.parse(atob(tokenPayload)) : null;
	if (decodedPayload) {
		event.locals.user = {
			name: decodedPayload?.username
		};
	}

	if (decodedPayload && event.route.id === '/login') {
		throw redirect(301, '/');
	} else if (decodedPayload || event.route.id === '/login') {
		return resolve(event);
	} else {
		throw redirect(301, '/login');
	}
}) satisfies Handle;
