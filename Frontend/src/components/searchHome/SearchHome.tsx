import SearchLogo from "../../assets/search.svg";
import "./searchHome.css";

const SearchHome = () => {
  return (
    <>
      <form>
        <input type="text" placeholder="What's on your to-do list?" />
        <button className="searchButton">
          <img src={SearchLogo} alt="search icon" />
        </button>
      </form>
    </>
  );
};

export default SearchHome;
