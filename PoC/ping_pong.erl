-module(ping_pong).
-export([start/0, run_go/0, loop/1]).

start() ->
  run_go(),
  {ok, Socket} = gen_tcp:connect("localhost", 8080, [binary, {packet, 0}, {active, false}]),
  loop(Socket).

run_go() ->
  % Command to start the Go application; ensure `main.go` is compiled and accessible in the PATH
  spawn(fun() -> os:cmd("go run main.go &") end).

loop(Socket) ->
  Message = "ping\n",
  io:format("Erlang: Sending ~p~n", [Message]),
  gen_tcp:send(Socket, Message),

  case gen_tcp:recv(Socket, 0) of
    {ok, Response} ->
      io:format("Erlang: Received ~p~n", [Response]),
      if Response =:= <<"pong\n">> ->
        loop(Socket);
        true ->
          io:format("Unexpected response: ~p~n", [Response]),
          gen_tcp:close(Socket)
      end;
    {error, Reason} ->
      io:format("Error: ~p~n", [Reason]),
      gen_tcp:close(Socket)
  end.
