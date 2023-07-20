import type { PageServerLoad } from './$types';
import { redirect } from '@sveltejs/kit';

export const load = (async ({ cookies, locals }) => {
	console.log('logout called');
	// We dont need our session token anymore
	cookies.delete('session');
	// Redirect to login
	throw redirect(301, '/login');
}) satisfies PageServerLoad;
