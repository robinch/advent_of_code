defmodule AOC2021.Day04 do
  def part_01(path) do
    {numbers_to_draw, boards} = read_inputs(path)

    {last_number, drawn_numbers, winning_board} =
      numbers_to_draw
      |> Enum.reduce_while({[], boards}, fn n, {drawn_numbers, boards} ->
        Enum.reduce_while(boards, {:ok, []}, fn board, {_, marked_boards} ->
          case mark(board, n) do
            {:bingo, marked_board} -> {:halt, {:bingo, marked_board}}
            {:ok, marked_board} -> {:cont, {:ok, [marked_board | marked_boards]}}
          end
        end)
        |> case do
          {:bingo, marked_board} -> {:halt, {n, [n | drawn_numbers], marked_board}}
          {:ok, marked_boards} -> {:cont, {[n | drawn_numbers], marked_boards}}
        end
      end)

    sum_of_numbers_not_marked = sum_of_not_marked(winning_board, drawn_numbers)
    sum_of_numbers_not_marked * String.to_integer(last_number)
  end

  def part_02(path) do
    {numbers_to_draw, boards} = read_inputs(path)

    {last_number, drawn_numbers, winning_board} =
      numbers_to_draw
      |> Enum.reduce_while({[], boards}, fn n, {drawn_numbers, boards} ->
        nr_of_boards = Enum.count(boards)

        Enum.reduce_while(boards, {:ok, [], nr_of_boards}, fn
          board, {_, marked_boards, nr_of_boards_left} ->
            case mark(board, n) do
              {:bingo, marked_board} ->
                cond do
                  nr_of_boards_left == 1 -> {:halt, {:last_board, marked_board}}
                  :otherwise -> {:cont, {:ok, marked_boards, nr_of_boards_left - 1}}
                end

              {:ok, marked_board} ->
                {:cont, {:ok, [marked_board | marked_boards], nr_of_boards_left}}
            end
        end)
        |> case do
          {:last_board, marked_board} -> {:halt, {n, [n | drawn_numbers], marked_board}}
          {:ok, marked_boards, _} -> {:cont, {[n | drawn_numbers], marked_boards}}
        end
      end)

    sum_of_numbers_not_marked = sum_of_not_marked(winning_board, drawn_numbers)

    sum_of_numbers_not_marked * String.to_integer(last_number)
  end

  def sum_of_not_marked(board, drawn_numbers) do
    board
    |> Map.get(:number_to_pos_map)
    |> Map.keys()
    |> Enum.filter(fn board_number -> board_number not in drawn_numbers end)
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum()
  end

  defp mark(board, n) do
    case board.number_to_pos_map[n] do
      nil ->
        {:ok, board}

      {row, col} ->
        marked_board =
          board
          |> mark_it(row)
          |> mark_it(col)

        row = get_in(marked_board, [:marked_map, row])
        col = get_in(marked_board, [:marked_map, col])

        cond do
          row == 5 || col == 5 -> {:bingo, marked_board}
          :otherwise -> {:ok, marked_board}
        end
    end
  end

  defp mark_it(board, row_or_col) do
    board
    |> update_in([:marked_map, row_or_col], fn
      nil -> 1
      occ -> occ + 1
    end)
  end

  defp read_inputs(path) do
    [drawn_numbers_input | board_inputs] =
      path
      |> File.read!()
      |> String.split("\n\n", trim: true)

    boards = Enum.map(board_inputs, &build_board/1)
    numbers_to_draw = drawn_numbers_input |> String.split(",")
    {numbers_to_draw, boards}
  end

  defp build_board(board) do
    number_to_pos_map =
      board
      |> String.split("\n", trim: true)
      |> Enum.flat_map(&String.split(&1, " ", trim: true))
      |> number_to_pos_map()

    %{number_to_pos_map: number_to_pos_map, marked_map: %{}}
  end

  defp number_to_pos_map(board) do
    board
    |> Enum.reduce({0, %{}}, fn n, {i, m} ->
      {i + 1, Map.put(m, n, {"r#{div(i, 5)}", "c#{rem(i, 5)}"})}
    end)
    |> elem(1)
  end
end
