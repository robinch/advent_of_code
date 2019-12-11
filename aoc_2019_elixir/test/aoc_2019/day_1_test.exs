defmodule Aoc2019.Day1Test do
  use ExUnit.Case

  alias Aoc2019.Day1

  describe "part 1" do
    test "mass 12, returns 2 fuel" do
      assert Day1.part_1(12) == 2
    end

    test "mass 14, returns 2 fuel" do
      assert Day1.part_1(14) == 2
    end

    test "mass 1969, returns 654 fuel" do
      assert Day1.part_1(1969) == 654
    end

    test "mass 100756, returns 33583 fuel" do
      assert Day1.part_1(100_756) == 33583
    end

    test "solution" do
      masses =
        File.read!("puzzle_inputs/day_1.txt")
        |> String.split("\n", trim: true)
        |> Enum.map(&String.to_integer/1)

      assert Day1.part_1(masses) == 3_406_527
    end
  end

  describe "part 2" do
    test "mass 14, returns 2 fuel" do
      assert Day1.part_2(12) == 2
    end

    test "mass 1969, returns 966 fuel" do
      assert Day1.part_2(1969) == 966
    end

    test "mass 100756, returns 50346 fuel" do
      assert Day1.part_2(100_756) == 50346
    end

    test "solution" do
      masses =
        File.read!("puzzle_inputs/day_1.txt")
        |> String.split("\n", trim: true)
        |> Enum.map(&String.to_integer/1)

      assert Day1.part_2(masses) == 5_106_932
    end
  end
end
