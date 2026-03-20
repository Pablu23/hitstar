const load = async ({ fetch, params }) => {
  const response = await fetch(`http://api.hitstar.xyz:8080/lobby/${params.id}/players/list`, {
    method: "GET",
    credentials: "include"
  });
  const playerList = await response.json();
  const players = playerList.map((player) => {
    return player.Name;
  });
  return {
    lobbyId: params.id,
    players
  };
};
export {
  load
};
