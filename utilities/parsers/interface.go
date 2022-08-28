package parsers

import (
	"os"
	"sort"
	"strings"
	"time"
)

func GetBlogPostByName(fileName string) (markdown string, meta *Meta) {
	rawData := newRawFileData(fileName)

	return newMarkdown(rawData), newMeta(rawData)
}

func GetBlogPostsSlice() []*Meta {
	var blogMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil
	}

	for _, f := range files {
		rawData := newRawFileData(strings.TrimSuffix(f.Name(), ".md"))

		meta := newMeta(rawData)

		blogMetaSlice = append(blogMetaSlice, meta)
	}

	// Sort by date and alphabetical
	sort.Slice(blogMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, blogMetaSlice[i].Date)
		t2, _ := time.Parse(format, blogMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return blogMetaSlice[i].Title < blogMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return blogMetaSlice
}

func GetSeriesSlice() []string {
	var seriesSlice []string

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil
	}

	for _, f := range files {
		rawData := newRawFileData(strings.TrimSuffix(f.Name(), ".md"))

		meta := newMeta(rawData)

		// check if series is "" and skip
		// go uses "" as a zero value for the string type
		// the "" will get pushed to the slice if not checked
		if meta.Series == "" {
			continue
		}

		seriesSlice = append(seriesSlice, meta.Series)
	}

	// Remove duplicate strings in slice
	var uniqueSeriesSlice []string
	for _, s := range seriesSlice {
		var same bool
		for _, a := range uniqueSeriesSlice {
			if a == s {
				same = true
				continue
			}
		}

		if !same {
			uniqueSeriesSlice = append(uniqueSeriesSlice, s)
		}
	}

	// Sort strings in alphabetical order
	sort.Strings(uniqueSeriesSlice)

	return uniqueSeriesSlice
}

func GetSliceBySeriesParam(series string) []*Meta {
	var seriesMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil
	}

	for _, f := range files {
		rawData := newRawFileData(strings.TrimSuffix(f.Name(), ".md"))

		meta := newMeta(rawData)

		if meta.Series == series {
			seriesMetaSlice = append(seriesMetaSlice, meta)
		}
	}

	// Sort by date and alphabetical
	sort.Slice(seriesMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, seriesMetaSlice[i].Date)
		t2, _ := time.Parse(format, seriesMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return seriesMetaSlice[i].Title < seriesMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return seriesMetaSlice
}

func GetSliceByCategoryParam(category string) []*Meta {
	var categoryMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil
	}

	for _, f := range files {
		rawData := newRawFileData(strings.TrimSuffix(f.Name(), ".md"))

		meta := newMeta(rawData)

		for _, t := range meta.Tags {
			if t == category {
				categoryMetaSlice = append(categoryMetaSlice, meta)
			}
		}
	}

	// Sort by date and alphabetical
	sort.Slice(categoryMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, categoryMetaSlice[i].Date)
		t2, _ := time.Parse(format, categoryMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return categoryMetaSlice[i].Title < categoryMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return categoryMetaSlice
}
