defmodule Aoc2019.Day3 do
  @spec part_1([binary], [binary]) :: integer
  def part_1(wire_a, wire_b) do
    closest_intersection_to_origin(wire_a, wire_b)
  end

  @spec part_2([binary], [binary]) :: integer
  def part_2(wire_a, wire_b) do
    shortest_wire_length_to_intersection(wire_a, wire_b)
  end

  defp closest_intersection_to_origin(wire_a, wire_b) do
    wire_a_map = wiring_map(wire_a)
    wire_b_map = wiring_map(wire_b)

    intersections(wire_a_map, wire_b_map)
    |> Enum.map(&manhattan_distance/1)
    |> Enum.min()
  end

  defp shortest_wire_length_to_intersection(wire_a, wire_b) do
    wire_a_map = wiring_map(wire_a)
    wire_b_map = wiring_map(wire_b)

    intersections(wire_a_map, wire_b_map)
    |> Enum.map(&wire_length_to_intersection(&1, wire_a_map, wire_b_map))
    |> Enum.min()
  end

  defp wiring_map(wire) do
    wire = Enum.map(wire, &to_tuple/1)

    Enum.reduce(wire, {%{}, {0, 0}, 0}, fn
      {direction, distance}, {map, current_position, traveled} ->
        draw_wiring_line(map, current_position, traveled, direction, distance)
    end)
    |> elem(0)
  end

  defp draw_wiring_line(map, current_position, traveled, direction, distance) do
    Enum.reduce(1..distance, {map, current_position, traveled}, fn
      _, {map, current_position, traveled} ->
        new_position = new_position(current_position, direction)
        traveled = traveled + 1
        {Map.put(map, new_position, traveled), new_position, traveled}
    end)
  end

  defp new_position({x, y}, "R"), do: {x + 1, y}
  defp new_position({x, y}, "L"), do: {x - 1, y}
  defp new_position({x, y}, "U"), do: {x, y + 1}
  defp new_position({x, y}, "D"), do: {x, y - 1}

  defp intersections(wire_a, wire_b) do
    {shorter_wire, longer_wire} =
      if Enum.count(wire_a) < Enum.count(wire_b) do
        {wire_a, wire_b}
      else
        {wire_b, wire_a}
      end

    Enum.filter(shorter_wire, fn {key, _val} -> Map.get(longer_wire, key) != nil end)
    |> Enum.map(&elem(&1, 0))
  end

  defp wire_length_to_intersection(intersection, wire_a, wire_b) do
    Map.get(wire_a, intersection) + Map.get(wire_b, intersection)
  end

  defp manhattan_distance({x, y}), do: abs(x) + abs(y)

  defp to_tuple(wiring) do
    {dir, dist} = String.split_at(wiring, 1)

    {dir, String.to_integer(dist)}
  end
end
