import { render, screen } from "@testing-library/react";
import "@testing-library/jest-dom";

import SearchHome from "./SearchHome";

describe("SearchHome", () => {
  it("renders SearchHome component", () => {
    render(<SearchHome />);
  });
  test("render input field", () => {
    render(<SearchHome />);
    expect(screen.getByRole("textbox")).toBeInTheDocument();
  });
  test("render search button", () => {
    render(<SearchHome />);
    expect(screen.getByRole("button")).toBeInTheDocument();
  });
});
