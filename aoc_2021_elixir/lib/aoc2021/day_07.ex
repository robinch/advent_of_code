defmodule AOC2021.Day07 do
  def part_01(path) do
    path
    |> File.read!()
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
    |> Enum.frequencies()
    |> magic(fn check_position, position, crabs_on_position ->
      abs(check_position - position) * crabs_on_position
    end)
  end

  def part_02(path) do
    path
    |> File.read!()
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
    |> Enum.frequencies()
    |> magic(fn check_position, position, crabs_on_position ->
      n = abs(check_position - position)

      n / 2 * (n + 1) * crabs_on_position
    end)
  end

  defp magic(map, fuel_fun) do
    positions = Map.keys(map)
    max = Enum.max(positions)

    min = Enum.min(positions)

    min..max
    |> Enum.map(fn check_position ->
      Enum.reduce(map, 0, fn {position, crabs_on_position}, fuel ->
        fuel + fuel_fun.(check_position, position, crabs_on_position)
      end)
    end)
    |> Enum.min()
    |> trunc()
  end
end
