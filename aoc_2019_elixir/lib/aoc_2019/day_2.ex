defmodule Aoc2019.Day2 do
  @spec part_1([integer]) :: [integer]
  def part_1(int_codes) do
    list_to_map(int_codes)
    |> run()
    |> map_to_list()
  end

  defp run(int_code_map, index \\ 0) do
    opcode = int_code_map[index]

    case opcode do
      99 ->
        int_code_map

      opcode ->
        {new_op_codes, new_index} = run_op(int_code_map, index, opcode)
        run(new_op_codes, new_index)
    end
  end

  defp run_op(int_code_map, index, opcode) do
    values = get_values(int_code_map, index)
    res = do_op(opcode, values)

    write_adr = Map.get(int_code_map, index + 3)
    {Map.put(int_code_map, write_adr, res), index + 4}
  end

  defp get_values(int_code_map, index) do
    read_adr_1 = Map.get(int_code_map, index + 1)
    read_adr_2 = Map.get(int_code_map, index + 2)
    val_1 = Map.get(int_code_map, read_adr_1)
    val_2 = Map.get(int_code_map, read_adr_2)

    {val_1, val_2}
  end

  defp do_op(1, {val_1, val_2}), do: val_1 + val_2
  defp do_op(2, {val_1, val_2}), do: val_1 * val_2

  defp list_to_map(list) do
    Enum.reduce(list, {%{}, 0}, fn
      v, {map, index} -> {Map.put(map, index, v), index + 1}
    end)
    |> elem(0)
  end

  defp map_to_list(map) do
    Enum.reduce((Enum.count(map) - 1)..0, [], fn i, acc ->
      [Map.get(map, i) | acc]
    end)
  end
end
