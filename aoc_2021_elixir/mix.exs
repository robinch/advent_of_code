defmodule Aoc2021.MixProject do
  use Mix.Project

  def project do
    [
      app: :aoc_2021,
      version: "0.1.0",
      elixir: "~> 1.12",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  def application do
    [
      extra_applications: [:logger]
    ]
  end

  defp deps do
    [
      {:mix_test_watch, "~> 1.1"}
    ]
  end
end
