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

  describe "Day 04" do
    test "part_1" do
      assert AOC2021.Day04.part_01("inputs/test/day_04.txt") == 4512
    end

    test "part_2" do
      assert AOC2021.Day04.part_02("inputs/test/day_04.txt") == 1924
    end
  end

  describe "Day 05" do
    # test "part_1" do
    #   assert AOC2021.Day05.part_01("inputs/test/day_05.txt") == 5
    # end

    # test "part_2" do
    #   assert AOC2021.Day05.part_02("inputs/test/day_05.txt") == 1924
    # end
  end

  describe "Day 06" do
    test "part_1" do
      assert AOC2021.Day06.part_01("inputs/test/day_06.txt") == 5934
    end

    test "part_2" do
      assert AOC2021.Day06.part_02("inputs/test/day_06.txt") == 26_984_457_539
    end
  end

  describe "Day 07" do
    test "part_1" do
      assert AOC2021.Day07.part_01("inputs/test/day_07.txt") == 37
    end

    test "part_2" do
      assert AOC2021.Day07.part_02("inputs/test/day_07.txt") == 168
    end
  end

  describe "Day 09" do
    # test "part_1" do
    #   assert AOC2021.Day09.part_01("inputs/test/day_09.txt") == 15
    # end

    # test "part_2" do
    #   assert AOC2021.Day09.part_02("inputs/test/day_09.txt") == 168
    # end
  end

  describe "Day 10" do
    test "part_1" do
      assert AOC2021.Day10.part_01("inputs/test/day_10.txt") == 26397
    end

    test "part_2" do
      assert AOC2021.Day10.part_02("inputs/day_10.txt") == 288_957
    end
  end

  describe "Day 20" do
    test "part_1" do
      assert AOC2021.Day20.part_01("inputs/test/day_20.txt") == 35
    end

    # test "part_2" do
    #   assert AOC2021.Day20.part_02("inputs/day_20.txt") == 288_957
    # end
  end
end
