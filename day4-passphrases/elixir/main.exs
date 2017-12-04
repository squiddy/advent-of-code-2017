defmodule Main do
    def check(passwords, seen \\ MapSet.new)
    def check([], _) do true end
    def check([h | t], seen) do
        case MapSet.member?(seen, h) do
            true -> false
            _ -> check(t, MapSet.put(seen, h))
        end
    end
end

File.stream!("../input.txt")
    |> Stream.map(&Main.check(String.split(&1)))
    |> Stream.filter(fn(x) -> x end)
    |> Enum.count
    |> IO.puts