// @ts-nocheck
import type { PageLoad } from './$types';

export const load = async ({ fetch, params }: Parameters<PageLoad>[0]) => {

  const response = await fetch(`http://api.hitstar.xyz:8080/lobby/${params.id}/players/list`, {
    method: 'GET',
    credentials: 'include'
  });

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const playerList: Array<any> = await response.json();
  const players = playerList.map((player) => {
      return player.Name;
    })

  return {
    lobbyId: params.id,
    players: players
  };
};
