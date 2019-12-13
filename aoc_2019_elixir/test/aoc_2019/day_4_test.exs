defmodule Aoc2019.Day4Test do
  use ExUnit.Case

  alias Aoc2019.Day4

  describe "part 1" do
    test "111111 meets the criteria" do
      assert Day4.part_1(111_111) == true
    end

    test "223450 does not meets the criteria" do
      assert Day4.part_1(223_450) == false
    end

    test "123789 does not meets the criteria" do
      assert Day4.part_1(123_789) == false
    end

    test "how many different passwords within 240920-789857 meets the criteria" do
      assert 240_920..789_857
             |> Enum.filter(&Day4.part_1/1)
             |> Enum.count() ==
               1154
    end
  end

  describe "part 2" do
    test "112233 meets the criteria" do
      assert Day4.part_2(112_233) == true
    end

    test "123444 meets the criteria" do
      assert Day4.part_2(123_444) == false
    end

    test "111122 meets the criteria" do
      assert Day4.part_2(111_122) == true
    end

    test "how many different passwords within 240920-789857 meets the criteria" do
      assert 240_920..789_857
             |> Enum.filter(&Day4.part_2/1)
             |> Enum.count() == 750
    end
  end
end
