defmodule AOC2021.Day05 do
  def part_01(path) do
    path
    |> File.read!()
    |> String.split("\n")
    |> Enum.map(&String.split(&1, " -> "))
  end
end
