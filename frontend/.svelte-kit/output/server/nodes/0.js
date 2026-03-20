import * as universal from '../entries/pages/_layout.ts.js';

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_layout.svelte.js')).default;
export { universal };
export const universal_id = "src/routes/+layout.ts";
export const imports = ["_app/immutable/nodes/0.DqmzXT2m.js","_app/immutable/chunks/DsnmJJEf.js","_app/immutable/chunks/Mxj_Vo4O.js","_app/immutable/chunks/Db9jlCb9.js","_app/immutable/chunks/CsvCdjdn.js"];
export const stylesheets = ["_app/immutable/assets/0.DNyseGkR.css"];
export const fonts = [];
