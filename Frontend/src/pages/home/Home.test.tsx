import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";

import Home from "./Home";

describe("Home", () => {
  it("renders Home", () => {
    render(<Home />);
  });

  test("returns home improvement made easy", () => {
    render(<Home />);
    expect(screen.getByText("Home,")).toBeInTheDocument();
    expect(screen.getByText("improvement,")).toBeInTheDocument();
    expect(screen.getByText("made easy")).toBeInTheDocument();
  });
});
