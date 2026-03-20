export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["robots.txt"]),
	mimeTypes: {".txt":"text/plain"},
	_: {
		client: {start:"_app/immutable/entry/start.CIN7fB44.js",app:"_app/immutable/entry/app.zvHQKp3W.js",imports:["_app/immutable/entry/start.CIN7fB44.js","_app/immutable/chunks/BYR8YxRs.js","_app/immutable/chunks/-Yrz3-Yz.js","_app/immutable/chunks/Mxj_Vo4O.js","_app/immutable/chunks/DyQ3jvN0.js","_app/immutable/chunks/Db9jlCb9.js","_app/immutable/chunks/D0iwhpLH.js","_app/immutable/entry/app.zvHQKp3W.js","_app/immutable/chunks/Mxj_Vo4O.js","_app/immutable/chunks/DyQ3jvN0.js","_app/immutable/chunks/DsnmJJEf.js","_app/immutable/chunks/-Yrz3-Yz.js","_app/immutable/chunks/Db9jlCb9.js","_app/immutable/chunks/DPbAKcnu.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js'))
		],
		remotes: {
			
		},
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/lobby/[id]",
				pattern: /^\/lobby\/([^/]+?)\/?$/,
				params: [{"name":"id","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/lobby/[id]/join",
				pattern: /^\/lobby\/([^/]+?)\/join\/?$/,
				params: [{"name":"id","optional":false,"rest":false,"chained":false}],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
