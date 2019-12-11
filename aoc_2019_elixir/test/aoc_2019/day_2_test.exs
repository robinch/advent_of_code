defmodule Aoc2019.Day2Test do
  use ExUnit.Case

  alias Aoc2019.Day2

  describe "part 1" do
    test "1,0,0,0,99 becomes 2,0,0,0,99" do
      assert Day2.part_1([1, 0, 0, 0, 99]) == [2, 0, 0, 0, 99]
    end

    test "2,3,0,3,99 becomes 2,3,0,6,99" do
      assert Day2.part_1([2, 3, 0, 3, 99]) == [2, 3, 0, 6, 99]
    end

    test "2,4,4,5,99,0 becomes 2,4,4,5,99,9801" do
      assert Day2.part_1([2, 4, 4, 5, 99, 0]) == [2, 4, 4, 5, 99, 9801]
    end

    test "1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99" do
      assert Day2.part_1([1, 1, 1, 4, 99, 5, 6, 0, 99]) == [30, 1, 1, 4, 2, 5, 6, 0, 99]
    end

    test "solution" do
      # I need replace these values
      # pos 1 with the value 12
      # pos 2 with the value 2
      int_codes =
        File.read!("puzzle_inputs/day_2.txt")
        |> String.split("\n", trim: true)
        |> hd()
        |> String.split(",", trim: true)
        |> Enum.map(&String.to_integer/1)

      assert Day2.part_1(int_codes, 12, 2) |> hd() == 2_782_414
    end
  end

  describe "part 2" do
  end
end
