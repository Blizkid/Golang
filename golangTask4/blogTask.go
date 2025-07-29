package main

import (
	"fmt"
	"gorm.io/gorm"
)

func RunBlog(db *gorm.DB) {
	// 自动建表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.AutoMigrate(&Comment{})

	//// 1. 创建用户
	//user1 := User{Name: "alice", Password: "password1", ArticleCount: 2}
	//user2 := User{Name: "bob", Password: "password2"}
	//db.Create(&user1)
	//db.Create(&user2)
	//
	//// 2. 创建帖子
	//post1 := Post{Title: "第一篇博客", Content: "内容1", UserID: user1.ID}
	//post2 := Post{Title: "第二篇博客", Content: "内容2", UserID: user1.ID}
	//post3 := Post{Title: "Bob的博客", Content: "内容3", UserID: user2.ID}

	//db.Create(&post1)
	//db.Create(&post2)
	//db.Create(&post3)
	//
	// 3. 创建评论
	comment1 := Comment{Content: "好文，学习了！", UserID: 1, PostID: 1}
	comment2 := Comment{Content: "期待更新", UserID: 2, PostID: 2}
	comment3 := Comment{Content: "棒棒哒", UserID: 3, PostID: 3}
	db.Create(&comment1)
	db.Create(&comment2)
	db.Create(&comment3)

	// 检查插入结果
	//var users []User
	//var posts []Post
	//var comments []Comment
	//db.Find(&users)
	//db.Find(&posts)
	//db.Find(&comments)
	//
	//fmt.Println("全部用户：", users)
	//fmt.Println("全部帖子：", posts)
	//fmt.Println("全部评论：", comments)

	//db.Preload("Posts.Comments").Where("id = ?", 3).First(&user1)
	//fmt.Printf("id=3用户的所有文章及其评论:\n")
	//for _, post := range user1.Posts {
	//	fmt.Printf("  文章: %s\n", post.Title)
	//	for _, comment := range post.Comments {
	//		fmt.Printf("    评论: %s\n", comment.Content)
	//	}
	//}

	db.Where("post_id = ?", 3).Unscoped().Delete(&Comment{PostID: 3})
	fmt.Println(db.RowsAffected)
}

type User struct {
	gorm.Model
	Name         string
	Password     string
	Posts        []Post
	ArticleCount int
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	Author        string
	Comments      []Comment
	UserID        uint // 外键
	CommentCount  int
	CommentStatus string

	AccountId User `gorm:"foreignKey:UserID;references:ID"`
}

type Comment struct {
	gorm.Model
	Content string
	Author  string
	PostID  uint
	UserID  uint // 外键

	RelatedPostID Post `gorm:"foreignKey:PostID;references:ID"`
	AccountId     User `gorm:"foreignKey:UserID;references:ID"`
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("article_count", gorm.Expr("article_count + ?", 1)).Error
	return
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete", c.PostID)
	var count int64
	err = tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 1 {
		err = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_count", 0).Update("comment_status", "无评论").Error
		if err != nil {
			return err
		}
	}
	return nil
}
