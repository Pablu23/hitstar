import { a as attr } from "../../chunks/attributes.js";
import { v as pop, t as push } from "../../chunks/index2.js";
import "@sveltejs/kit/internal";
import "../../chunks/exports.js";
import "../../chunks/utils.js";
import "../../chunks/state.svelte.js";
function _page($$payload, $$props) {
  push();
  let lobbyId = "";
  $$payload.out.push(`<h1>Welcome to SvelteKit</h1> <p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p> <a href="http://api.hitstar.xyz:8080/login">Login</a> <button>Create Lobby</button> <input${attr("value", lobbyId)}/> `);
  {
    $$payload.out.push("<!--[!-->");
  }
  $$payload.out.push(`<!--]-->`);
  pop();
}
export {
  _page as default
};
