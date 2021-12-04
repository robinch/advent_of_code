defmodule AOC2021Test do
  use ExUnit.Case

  describe "Day 01" do
    test "part_1" do
      assert AOC2021.Day01.part_01("inputs/test/day_01.txt") == 7
    end

    test "part_2" do
      assert AOC2021.Day01.part_02("inputs/test/day_01.txt") == 5
    end
  end

  describe "Day 02" do
    test "part_1" do
      assert AOC2021.Day02.part_01("inputs/test/day_02.txt") == 150
    end

    test "part_2" do
      assert AOC2021.Day02.part_02("inputs/test/day_02.txt") == 900
    end
  end

  describe "Day 03" do
    test "part_1" do
      assert AOC2021.Day03.part_01("inputs/test/day_03.txt") == 198
    end

    test "part_2" do
      assert AOC2021.Day03.part_02("inputs/test/day_03.txt") == 230
    end
  end
end
