defmodule AOC2021.Day02 do
  def part_01(path) do
    {horizontal, depth} =
      read_input(path)
      |> Enum.reduce({0, 0}, fn
        ["down", distance], {horizontal, depth} -> {horizontal, depth + distance}
        ["up", distance], {horizontal, depth} -> {horizontal, depth - distance}
        ["forward", distance], {horizontal, depth} -> {horizontal + distance, depth}
      end)

    horizontal * depth
  end

  def part_02(path) do
    {horizontal, depth, _aim} =
      read_input(path)
      |> Enum.reduce({0, 0, 0}, fn
        ["down", distance], {horizontal, depth, aim} ->
          {horizontal, depth, aim + distance}

        ["up", distance], {horizontal, depth, aim} ->
          {horizontal, depth, aim - distance}

        ["forward", distance], {horizontal, depth, aim} ->
          {horizontal + distance, depth + distance * aim, aim}
      end)

    I
    horizontal * depth
  end

  defp read_input(path) do
    path
    |> File.stream!()
    |> Stream.map(&String.replace(&1, "\n", ""))
    |> Stream.map(&String.split(&1, " "))
    |> Stream.map(fn [direction, distance] -> [direction, String.to_integer(distance)] end)
  end
end
