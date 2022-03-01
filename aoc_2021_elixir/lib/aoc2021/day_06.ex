defmodule AOC2021.Day06 do
  def part_01(path) do
    path
    |> input()
    |> Enum.frequencies()
    |> fast_calculate(80)
  end

  def part_02(path) do
    path
    |> input()
    |> Enum.frequencies()
    |> fast_calculate(256)
  end

  defp fast_calculate(current_state, 0) do
    current_state
    |> Map.values()
    |> Enum.sum()
  end

  defp fast_calculate(current_state, days_left) do
    current_state
    |> Enum.reduce(%{}, fn
      {0, fish}, acc ->
        acc
        |> Map.put(8, fish)
        |> Map.update(6, fish, &(&1 + fish))

      {days_until_spawn, fish}, acc ->
        Map.update(acc, days_until_spawn - 1, fish, &(&1 + fish))
    end)
    |> fast_calculate(days_left - 1)
  end

  defp calculate_lanternfish(current_state, days_left),
    do: calculate_lanternfish(current_state, [], days_left)

  defp calculate_lanternfish(current_state, _new_state, 0), do: Enum.count(current_state)

  defp calculate_lanternfish([], new_state, days_left),
    do: calculate_lanternfish(new_state, [], days_left - 1)

  defp calculate_lanternfish([0 | rest], new_state, days_left) do
    calculate_lanternfish(rest, [8, 6 | new_state], days_left)
  end

  defp calculate_lanternfish([d | rest], new_state, days_left) do
    calculate_lanternfish(rest, [d - 1 | new_state], days_left)
  end

  def input(path) do
    path
    |> File.read!()
    |> String.split(",", trim: true)
    |> Enum.map(&String.to_integer/1)
  end
end
