defmodule Day3.Part2 do
  def solve() do
    File.stream!("input.txt")
    |> Stream.map(&String.replace(&1, "\n", ""))
    |> to_cleaned_stream()
    |> to_collision_map()
    |> find_ids_without_any_collisions()
  end

  def clean_id(id_string) do
    String.slice(id_string, 1, String.length(id_string) - 1)
  end

  def clean_position(position_string) do
    position_string
    |> String.slice(0, String.length(position_string) - 1)
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
    |> List.to_tuple()
  end

  def clean_dimensions(dimension_string) do
    dimension_string
    |> String.split("x")
    |> Enum.map(&String.to_integer/1)
    |> List.to_tuple()
  end

  def to_cleaned_stream(uncleaned_stream) do
    Stream.map(uncleaned_stream, fn claim ->
      [id, _, position_string, dimension_string] = String.split(claim)
      {clean_id(id), clean_position(position_string), clean_dimensions(dimension_string)}
    end)
  end

  def to_collision_map(collision_map) do
    Enum.reduce(collision_map, %{}, fn {id, {start_x, start_y}, {width, height}}, map ->
      Enum.reduce(start_x..(start_x + width - 1), map, fn x, acc ->
        Enum.reduce(start_y..(start_y + height - 1), acc, fn y, map ->
          Map.update(map, {x, y}, {:no_collision, [id]}, fn {_collision_status, id_list} ->
            {:collision, [id | id_list]}
          end)
        end)
      end)
    end)
  end

  def find_ids_without_any_collisions(collision_map) do
    {collision_set, no_collision_set} =
      Enum.reduce(
        collision_map,
        {_collision_set = MapSet.new(), _no_collision_set = MapSet.new()},
        fn
          {_index, {:collision, ids}}, {collision_set, no_collision_set} ->
            collision_set =
              Enum.reduce(ids, collision_set, fn id, set ->
                MapSet.put(set, id)
              end)

            {collision_set, no_collision_set}

          {_index, {:no_collision, [id]}}, {collision_set, no_collision_set} ->
            no_collision_set = MapSet.put(no_collision_set, id)

            {collision_set, no_collision_set}
        end
      )

    MapSet.difference(no_collision_set, collision_set)
    |> MapSet.to_list()
  end
end
