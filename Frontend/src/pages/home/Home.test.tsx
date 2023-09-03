import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";

import Home from "./Home";

describe("Home", () => {
  it("renders Home", () => {
    render(<Home />);
  });

  test("shows hero heading", () => {
    render(<Home />);
    expect(screen.getByText("Home,")).toBeInTheDocument();
    expect(screen.getByText("improvement,")).toBeInTheDocument();
    expect(screen.getByText("made easy")).toBeInTheDocument();
  });

  test("shows sub heading", () => {
    render(<Home />);
    expect(screen.getByText(/^Try searching/)).toBeInTheDocument();
    expect(screen.getByText(/-*Plumber-*/)).toBeInTheDocument();
    expect(screen.getByText(/-*Handyman-*/)).toBeInTheDocument();
  });
});
