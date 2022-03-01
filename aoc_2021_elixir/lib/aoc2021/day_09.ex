defmodule AOC2021.Day09 do
  def part_01(path) do
    input =
      path
      |> File.read!()
      |> String.split("\n")
      |> Enum.map(fn s -> String.codepoints(s) |> Enum.map(&String.to_integer/1) end)
      |> Enum.map(&List.to_tuple/1)
      |> List.to_tuple()
      |> IO.inspect()

    get(input, 0, 1)
  end

  defp get(tuple, row, col) do
    tuple
    |> elem(row)
    |> elem(col)
  end
end
