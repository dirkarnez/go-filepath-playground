					dir := fmt.Sprintf(`%s\%s\%s\Lesson %d\%s\*`, prefix, rootDir, levels[j].Name, lessonNumber, folderName)
					filePaths, err := filepath.Glob(dir)
					if err != nil {
						log.Fatal(err)
					}

					filePaths = lo.Filter(filePaths, func(item string, index int) bool {
						return item != ".gitkeep"
					})
