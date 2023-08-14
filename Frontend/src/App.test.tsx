import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";

import App from "./App";

describe("App", () => {
  it("renders headline", () => {
    render(<App />);
    expect(screen.getByText("Vite + React")).toBeInTheDocument();
    screen.getByRole("button");
  });
});
