remove_ds_store:
	find . -name .DS_Store -print0 | xargs -0 git rm -f --ignore-unmatch
