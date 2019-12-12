defmodule Aoc2019.Day3Test do
  use ExUnit.Case
  alias Aoc2019.Day3

  test "R8,U5,L5,D3 and U7,R6,D4,L4 gives 6" do
    assert Day3.part_1(["R8", "U5", "L5", "D3"], ["U7", "R6", "D4", "L4"]) == 6
  end

  test "R75,D30,R83,U83,L12,D49,R71,U7,L72 and U62,R66,U55,R34,D71,R55,D58,R83 gives 159" do
    assert Day3.part_1(
             ["R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"],
             ["U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"]
           ) == 159
  end

  test "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51 and U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 gives 135" do
    assert Day3.part_1(
             ["R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"],
             ["U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"]
           ) == 135
  end

  test "solution" do
    [wire_a, wire_b] =
      File.read!("puzzle_inputs/day_3.txt")
      |> String.split("\n", trim: true)
      |> Enum.map(&String.split(&1, ","))

    assert Day3.part_1(wire_a, wire_b) == 207
  end
end
