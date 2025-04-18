import { goto } from '$app/navigation';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ parent }) => {
	const { user } = await parent();

	if (!user) {
		return;
	}

	if (user) {
		goto('/dashboard');
		return;
	}
};
