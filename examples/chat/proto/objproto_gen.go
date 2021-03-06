// Generated by github.com/davyxu/cellnet/objprotogen
// DO NOT EDIT!
package chatproto

import (
	"github.com/davyxu/cellnet"
	"reflect"
	"fmt"
)

func (self *PublicChatREQ) String() string { return fmt.Sprintf("%+v", *self) }
func (self *PublicChatACK) String() string { return fmt.Sprintf("%+v", *self) }
func (self *RenameACK) String() string     { return fmt.Sprintf("%+v", *self) }
func (self *GetNameREQ) String() string    { return fmt.Sprintf("%+v", *self) }
func (self *GetNameACK) String() string    { return fmt.Sprintf("%+v", *self) }

func init() {

	cellnet.RegisterMessageMeta("binary", "chatproto.PublicChatREQ", reflect.TypeOf((*PublicChatREQ)(nil)).Elem(), 1094803174)
	cellnet.RegisterMessageMeta("binary", "chatproto.PublicChatACK", reflect.TypeOf((*PublicChatACK)(nil)).Elem(), 706454148)
	cellnet.RegisterMessageMeta("binary", "chatproto.RenameACK", reflect.TypeOf((*RenameACK)(nil)).Elem(), 1608756819)
	cellnet.RegisterMessageMeta("binary", "chatproto.GetNameREQ", reflect.TypeOf((*GetNameREQ)(nil)).Elem(), 1208313923)
	cellnet.RegisterMessageMeta("binary", "chatproto.GetNameACK", reflect.TypeOf((*GetNameACK)(nil)).Elem(), 593466401)
}
