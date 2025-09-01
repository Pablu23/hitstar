<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

<a href="http://api.hitstar.xyz:8080/login">Login</a>

<button
	on:click={async (e) => {
		const response = await fetch('http://api.hitstar.xyz:8080/createLobby', {
			method: 'POST',
			credentials: 'include'
		});

		console.log(await response.text());
	}}>Create Lobby</button
>
