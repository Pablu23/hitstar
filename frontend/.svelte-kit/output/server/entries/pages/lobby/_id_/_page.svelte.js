import { y as ensure_array_like, v as pop, t as push } from "../../../../chunks/index2.js";
import { e as escape_html } from "../../../../chunks/escaping.js";
function _page($$payload, $$props) {
  push();
  let { data } = $$props;
  let players = data.players;
  const each_array = ensure_array_like(players);
  $$payload.out.push(`<h1>${escape_html(data.lobbyId)}</h1> <p>You are</p> <ul><!--[-->`);
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let player = each_array[$$index];
    $$payload.out.push(`<li>${escape_html(player)}</li>`);
  }
  $$payload.out.push(`<!--]--></ul>`);
  pop();
}
export {
  _page as default
};
