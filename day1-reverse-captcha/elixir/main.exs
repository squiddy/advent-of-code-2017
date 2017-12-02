defmodule Main do
    def reverse_captcha(input, lookahead \\ 1) do
        numbers = String.codepoints(input) |> Enum.map(&String.to_integer(&1))
        numbers_lookahead = Stream.drop(Stream.cycle(numbers), lookahead)
        pairs = Stream.zip([numbers, numbers_lookahead])

        Enum.reduce(pairs, 0,
            fn(pair, acc) ->
                case pair do
                    {x, x} -> x + acc
                    _ -> acc
                end
            end)
    end

    def reverse_captcha2(input) do
        reverse_captcha(input, String.length(input) / 2)
    end
end

IO.puts Main.reverse_captcha2(hd System.argv)