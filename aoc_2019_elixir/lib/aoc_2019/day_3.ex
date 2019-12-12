defmodule Aoc2019.Day3 do
  @spec part_1([binary], [binary]) :: integer
  def part_1(wire_a, wire_b) do
    closest_intersection_to_origin(wire_a, wire_b)
  end

  defp closest_intersection_to_origin(wire_a, wire_b) do
    wire_a_map = wiring_map(wire_a)
    wire_b_map = wiring_map(wire_b)

    MapSet.intersection(wire_a_map, wire_b_map)
    |> Enum.map(&manhattan_distance/1)
    |> Enum.min()
  end

  defp wiring_map(wire) do
    wire = Enum.map(wire, &to_tuple/1)

    Enum.reduce(wire, {MapSet.new(), {0, 0}}, fn {direction, distance}, {map, current_position} ->
      draw_wiring_line(map, current_position, direction, distance)
    end)
    |> elem(0)
  end

  defp draw_wiring_line(map, current_position, direction, distance) do
    Enum.reduce(1..distance, {map, current_position}, fn
      _, {map, current_position} ->
        new_position = new_position(current_position, direction)
        {MapSet.put(map, new_position), new_position}
    end)
  end

  defp new_position({x, y}, "R"), do: {x + 1, y}
  defp new_position({x, y}, "L"), do: {x - 1, y}
  defp new_position({x, y}, "U"), do: {x, y + 1}
  defp new_position({x, y}, "D"), do: {x, y - 1}

  defp manhattan_distance({x, y}), do: abs(x) + abs(y)

  defp to_tuple(wiring) do
    {dir, dist} = String.split_at(wiring, 1)

    {dir, String.to_integer(dist)}
  end
end
