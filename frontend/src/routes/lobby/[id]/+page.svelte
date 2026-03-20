<script lang="ts">
	import type { PageProps } from './$types';

	let { data }: PageProps = $props();
  let players = $state(data.players);
</script>

<h1>{data.lobbyId}</h1>

<p> You are </p>

<ul>
  {#each players as player (player)}
    <li>
      {player}
    </li>
  {/each}
</ul>
