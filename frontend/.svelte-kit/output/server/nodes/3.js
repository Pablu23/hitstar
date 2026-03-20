import * as universal from '../entries/pages/lobby/_id_/_page.ts.js';

export const index = 3;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/lobby/_id_/_page.svelte.js')).default;
export { universal };
export const universal_id = "src/routes/lobby/[id]/+page.ts";
export const imports = ["_app/immutable/nodes/3.Do0AfRbc.js","_app/immutable/chunks/DsnmJJEf.js","_app/immutable/chunks/Mxj_Vo4O.js","_app/immutable/chunks/DyQ3jvN0.js"];
export const stylesheets = [];
export const fonts = [];
