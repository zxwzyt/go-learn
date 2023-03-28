package training

import (
	"log"
	"zxw-go/common"
)

func MakeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := common.ImageFile(f); err != nil {
			log.Println(err)
		}

		// go common.ImageFile(f)
	}
}

func MakeThumbnails2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(t string) {
			common.ImageFile(t)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func MakeThumbnails3(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			if _, err := common.ImageFile(f); err != nil {
				errors <- err
			}
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func MakeThumbnails4(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = common.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}
