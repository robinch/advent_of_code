defmodule AOC2021.Day01 do
  def part_01(path) do
    File.stream!(path)
    |> Stream.map(&String.replace(&1, "\n", ""))
    |> Stream.map(&String.to_integer/1)
    |> Enum.reduce({-1, -1}, fn
      current_depth, {previous_depth, nr_of_increases} when current_depth > previous_depth ->
        {current_depth, nr_of_increases + 1}

      current_depth, {_, nr_of_increases} ->
        {current_depth, nr_of_increases}
    end)
    |> elem(1)
  end

  def part_02(path) do
    File.read!(path)
    |> String.split("\n", trim: true)
    |> Enum.map(&String.to_integer/1)
    |> three_measurement_window()
  end

  defp three_measurement_window(list), do: three_measurement_window(list, 0)
  defp three_measurement_window([_, _, _ | []], acc), do: acc

  defp three_measurement_window([elem1, _, _, elem4 | _] = list, acc) when elem4 > elem1,
    do: three_measurement_window(tl(list), acc + 1)

  defp three_measurement_window(list, acc), do: three_measurement_window(tl(list), acc)
end
