import Searchicon from "../../assets/search.svg";
import "./searchHome.css";

const SearchHome = ({ onSearchClick }) => {
  return (
    <>
      <form>
        <input
          type="text"
          placeholder="What's on your to-do list?"
          onClick={onSearchClick}
        />
        <button className="searchButton" onClick={onSearchClick}>
          <img src={Searchicon} alt="search icon" />
        </button>
      </form>
    </>
  );
};

export default SearchHome;
