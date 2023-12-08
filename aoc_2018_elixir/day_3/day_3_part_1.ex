defmodule Day3.Part1 do
  def solve() do
    File.stream!("input.txt")
    |> Stream.map(&String.replace(&1, "\n", ""))
    |> to_cleaned_stream()
    |> to_collision_map()
    |> count_collisions()
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
      [_, _, position_string, dimension_string] = String.split(claim)
      {clean_position(position_string), clean_dimensions(dimension_string)}
    end)
  end

  def to_collision_map(collision_map) do
    Enum.reduce(collision_map, %{}, fn {{start_x, start_y}, {width, height}}, map ->
      Enum.reduce(start_x..(start_x + width - 1), map, fn x, acc ->
        Enum.reduce(start_y..(start_y + height - 1), acc, fn y, map ->
          Map.update(map, {x, y}, :claimed, fn _claimed -> :collision end)
        end)
      end)
    end)
  end

  def count_collisions(collision_map) do
    Enum.reduce(collision_map, 0, fn
      {_index, :collision}, nr_of_claimed -> nr_of_claimed + 1
      {_index, _status}, nr_of_claimed -> nr_of_claimed
    end)
  end
end

ExUnit.start()

defmodule Day3.Part1Test do
  use ExUnit.Case

  test "clean id" do
    assert Day3.Part1.clean_id("#1234") == "1234"
  end

  test "clean_position" do
    assert Day3.Part1.clean_position("935,649:") == {935, 649}
  end

  test "clean dimensions" do
    assert Day3.Part1.clean_dimensions("12x34") == {12, 34}
  end
end

Day3.Part1.solve()
|> IO.inspect(label: "results")
