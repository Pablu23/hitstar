import { error, redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types"


export const load: PageLoad = async ({ fetch, params }) => {
  const response = await fetch(`http://api.hitstar.xyz:8080/lobby/${params.id}/join`, {
    method: 'POST',
    credentials: 'include'
  });

  if (!response.ok) {
    error(401, "Unauthorized");
  }

  redirect(307, `/lobby/${params.id}`);
}
